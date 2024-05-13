use super::config::*;
use super::options::SetOptions;
use super::{ExpirationEntry, Item, NestedMap, NestedValue};
use std::collections::VecDeque;
use std::sync::atomic::Ordering;
use std::time::SystemTime;

impl NestedMap {
    pub fn set(&mut self, keys: &str, value: &[u8], options: Option<SetOptions>) {
        let options = options.unwrap_or_default();
        let mut current_map = &mut self.data;

        // Traverse to the appropriate node
        for key in keys.split(DELIMITER) {
            current_map = current_map
                .entry(key.to_string())
                .or_insert_with(|| NestedValue::Map(NestedMap::new(self.max_history)))
                .as_map_mut();
        }

        // Access or create the items list at the final key under VALUE_KEY
        let items = current_map
            .entry(VALUE_KEY.to_string())
            .or_insert_with(|| NestedValue::Items(VecDeque::new()));

        if let NestedValue::Items(items) = items {
            let id = self.id_counter.fetch_add(1, Ordering::Relaxed);
            let new_item = Item {
                key: keys.to_string(),
                value: value.to_vec(),
                timestamp: SystemTime::now(),
                id,
            };

            let length: usize = items.len();

            if !options.preserve_history {
                if length > 0 {
                    items[0] = new_item;
                } else {
                    items.insert(0, new_item);
                }

                return;
            }

            // Prepend new item to the list to keep the newest items at the start
            if length >= self.max_history {
                items.pop_back(); // Remove the oldest item if we exceed the max history
            }
            items.push_front(new_item); // Insert new item at the start of the list

            let expires_at = SystemTime::now() + options.ttl;
            let expiration_entry = ExpirationEntry {
                expires_at,
                id,
                keys: keys.to_string(),
            };

            self.exp_heap.lock().unwrap().push(expiration_entry);
        }
    }
}

mod tests {
    use super::*;
    use crate::nestedmap::options::*;
    use crate::nestedmap::test_helpers::*;

    #[test]
    fn test_set() {
        //let mut nm = NestedMap::new(1);

        let test_cases = vec![
            TestCase {
                name: "Test depth 1",
                setup: Box::new(|nm| {
                    nm.set("a", b"the value a", None);
                }),
                search_keys: "a".to_string(),
                expected: vec![create_item("a", b"the value a")],
                max_history: 1,
            },
            TestCase {
                name: "Test depth 3",
                setup: Box::new(|nm| {
                    nm.set("a.b.c", b"the value abc", None);
                }),
                search_keys: "a.b.c".to_string(),
                expected: vec![create_item("a.b.c", b"the value abc")],
                max_history: 1,
            },
            TestCase {
                name: "Test depth 6",
                setup: Box::new(|nm| {
                    nm.set("a.b.c.d.e.f", b"the value abcdef", None);
                }),
                search_keys: "a.b.c.d.e.f".to_string(),
                expected: vec![create_item("a.b.c.d.e.f", b"the value abcdef")],
                max_history: 1,
            },
        ];

        set_tests(test_cases)
    }

    #[test]
    fn test_set_without_history() {
        let test_cases = vec![TestCase {
            name: "Test without history option",
            setup: Box::new(|nm| {
                for i in 1..=7 {
                    nm.set(
                        "a.b.c.d",
                        &format!("value{}", i).into_bytes(),
                        Some(SetOptions::new().preserve_history(false)),
                    );
                }
            }),
            search_keys: "a.b.c.d".to_string(),
            expected: vec![create_item("a.b.c.d", b"value7")],
            max_history: 5,
        }];

        set_tests(test_cases)
    }

    #[test]
    fn test_set_history() {
        let test_cases = vec![
            TestCase {
                name: "Test more than max_history values",
                setup: Box::new(|nm| {
                    for i in 1..=7 {
                        nm.set(
                            "a.b.c.d",
                            &format!("value{}", i).into_bytes(),
                            Some(SetOptions::new().preserve_history(true)),
                        );
                    }
                }),
                search_keys: "a.b.c.d".to_string(),
                expected: vec![
                    create_item("a.b.c.d", b"value7"),
                    create_item("a.b.c.d", b"value6"),
                    create_item("a.b.c.d", b"value5"),
                    create_item("a.b.c.d", b"value4"),
                    create_item("a.b.c.d", b"value3"),
                ],
                max_history: 5,
            },
            TestCase {
                name: "Test less than max_history values",
                setup: Box::new(|nm| {
                    for i in 1..=3 {
                        nm.set(
                            "a.b.c.d",
                            &format!("value{}", i).into_bytes(),
                            Some(SetOptions::new().preserve_history(true)),
                        );
                    }
                }),
                search_keys: "a.b.c.d".to_string(),
                expected: vec![
                    create_item("a.b.c.d", b"value3"),
                    create_item("a.b.c.d", b"value2"),
                    create_item("a.b.c.d", b"value1"),
                ],
                max_history: 5,
            },
            TestCase {
                name: "Test exactly max_history values",
                setup: Box::new(|nm| {
                    for i in 1..=5 {
                        nm.set(
                            "a.b.c.d",
                            &format!("value{}", i).into_bytes(),
                            Some(SetOptions::new().preserve_history(true)),
                        );
                    }
                }),
                search_keys: "a.b.c.d".to_string(),
                expected: vec![
                    create_item("a.b.c.d", b"value5"),
                    create_item("a.b.c.d", b"value4"),
                    create_item("a.b.c.d", b"value3"),
                    create_item("a.b.c.d", b"value2"),
                    create_item("a.b.c.d", b"value1"),
                ],
                max_history: 5,
            },
        ];

        set_tests(test_cases)
    }

    #[test]
    fn test_set_mixed_history() {
        let test_cases = vec![TestCase {
            name: "Test more than max_history values",
            setup: Box::new(|nm| {
                nm.set(
                    "a.b.c.d",
                    b"value1",
                    Some(SetOptions::new().preserve_history(true)),
                );
                nm.set(
                    "a.b.c.d",
                    b"value2",
                    Some(SetOptions::new().preserve_history(true)),
                );
                nm.set(
                    "a.b.c.d",
                    b"value3",
                    Some(SetOptions::new().preserve_history(true)),
                );
                nm.set(
                    "a.b.c.d",
                    b"value4",
                    Some(SetOptions::new().preserve_history(false)),
                );
                nm.set(
                    "a.b.c.d",
                    b"value5",
                    Some(SetOptions::new().preserve_history(true)),
                );
            }),
            search_keys: "a.b.c.d".to_string(),
            expected: vec![
                create_item("a.b.c.d", b"value5"),
                create_item("a.b.c.d", b"value4"),
                create_item("a.b.c.d", b"value2"),
                create_item("a.b.c.d", b"value1"),
            ],
            max_history: 5,
        }];

        set_tests(test_cases)
    }

    fn set_tests(test_cases: Vec<TestCase>) {
        for test in test_cases {
            let mut nm = NestedMap::new(test.max_history);
            (test.setup)(&mut nm);

            let results = nm.query(
                &test.search_keys,
                Some(GetOptions::new().history_count(test.max_history)),
            );
            assert_eq!(results.len(), test.expected.len());
            for (i, v) in results.iter().enumerate() {
                assert!(items_equal(v, &test.expected[i]));
            }
        }
    }
}
