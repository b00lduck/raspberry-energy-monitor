

INSERT INTO counter (id, name, unit) VALUES (1, 'Erdgas', 'm³');
INSERT INTO counter (id, name, unit) VALUES (2, 'Strom 1', 'kWh');
INSERT INTO counter (id, name, unit) VALUES (3, 'Strom 2', 'kWh');
INSERT INTO counter (id, name, unit) VALUES (4, 'Wasser', 'm³');

UPDATE hibernate_sequences SET sequence_next_hi_value = 5 WHERE sequence_name = 'counter';
