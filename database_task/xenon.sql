CREATE DATABASE xenon;
USE xenon;
SELECT * FROM events;
SELECT * FROM events_data;
SELECT * FROM selector_data;
SELECT * FROM events_data;
SELECT serial_number FROM events_data WHERE serial_number='587860';
SELECT * FROM events_data WHERE serial_number='587860';
SELECT SUM(TABLE_ROWS) 
     FROM INFORMATION_SCHEMA.TABLES 
     WHERE TABLE_SCHEMA = '{events}';
SELECT COUNT(*) FROM events;
SELECT COUNT(*) FROM events_data;
SELECT 
    serial_number, 
    COUNT(serial_number)
FROM
    events_data
GROUP BY serial_number
HAVING COUNT(serial_number) > 1;

SELECT 
    serial_number, COUNT(serial_number),
    event_timestamp,  COUNT(event_timestamp)
FROM
    events_data
GROUP BY 
    serial_number , 
    event_timestamp 
HAVING  
COUNT(serial_number) > 1
    AND COUNT(event_timestamp) > 1;

SELECT MIN(event_date), MAX(event_date) FROM events_data;
SELECT * FROM events_data ORDER BY serial_number ASC, event_timestamp DESC;
DESC events_data;
SELECT events.serial_number,selector_data.serial_number
FROM events 
INNER JOIN selector_data
ON events.serial_number = selector_data.serial_number;










