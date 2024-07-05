----- Procedure -----
create procedure insert_book(character varying,character varying,character varying)
language 'plpgsql'
as $$
begin 
	insert into public.todo_items (title, description, status) values ($1,$2,$3);
	commit;
end;
$$;

call insert_book('a12','test_procedure1','available')

select * from public.todo_items

----- Trigger -----
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Tạo trigger
CREATE TRIGGER update_timestamp
BEFORE UPDATE ON public.todo_items
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- Check update
UPDATE public.todo_items 
SET title = 'a1_updated' 
WHERE id = 21;

select * from public.todo_items



----- INDEX -----
-- Tạo index trên cột Title
CREATE INDEX idx_todo_items_title ON todo_items(title);

-- Tạo index trên cột Description
CREATE INDEX idx_todo_items_description ON todo_items(description);

-- Tạo index trên cột Status
CREATE INDEX idx_todo_items_status ON todo_items(status);

-- Tạo index kết hợp trên các cột Title và Status
CREATE INDEX idx_todo_items_title_status ON todo_items(title, status);

select * from public.todo_items where title = 'a123456' and status = 'updated'
select * from public.todo_items where title = 'a123456'
select * from public.todo_items where status = 'updated'
select * from public.todo_items where status = 'available'


explain analyze select * from public.todo_items where status = 'available' and title = 'a123456'
explain analyze select * from public.todo_items where title = 'a123456' and status = 'available'

explain analyze select * from public.todo_items where title = 'a123456'
explain analyze select * from public.todo_items where status = 'updated'
	
SELECT * FROM public.todo_items WHERE title = 'title_123456' AND status = 'available';
-- Truy vấn kết hợp title và status
EXPLAIN ANALYZE SELECT * FROM public.todo_items WHERE title = 'title_123456' AND status = 'available';

	
drop index idx_todo_items_title_status
drop index idx_todo_items_description
drop index idx_todo_items_title
drop index idx_todo_items_status






-- Thêm dữ liệu mẫu
DO $$
BEGIN
    FOR i IN 1..5000000 LOOP
        INSERT INTO public.todo_items (title, description, status)
        VALUES (
            'title_' || i,
            'description_' || i,
            CASE 
                WHEN i % 2 = 0 THEN 'available' 
                ELSE 'unavailable' 
            END
        );
    END LOOP;
END $$;




----- Procedure -----
CREATE OR REPLACE PROCEDURE insert_book_test(
    IN p_title CHARACTER VARYING,
    IN p_description CHARACTER VARYING,
    IN p_status CHARACTER VARYING,
    OUT o_id INTEGER,
    OUT o_title CHARACTER VARYING,
    OUT o_description CHARACTER VARYING,
    OUT o_status CHARACTER VARYING,
    OUT o_count INTEGER
)
LANGUAGE plpgsql
AS $$
BEGIN 
    INSERT INTO public.todo_items (title, description, status) 
    VALUES (p_title, p_description, p_status)
    RETURNING id, title, description, status 
    INTO o_id, o_title, o_description, o_status;
    
    COMMIT;
    SELECT COUNT(*) INTO o_count FROM public.todo_items;
END;
$$;

DO $$
DECLARE
    v_id INTEGER;
    v_title CHARACTER VARYING;
    v_description CHARACTER VARYING;
    v_status CHARACTER VARYING;
    v_count INTEGER; 
BEGIN
    CALL insert_book_test('a12345678', 'test_procedure1', 'available', v_id, v_title, v_description, v_status, v_count);
    
    RAISE NOTICE 'Inserted book: ID %, Title %, Description %, Status %', v_id, v_title, v_description, v_status;
    RAISE NOTICE 'Total number of records: %', v_count; 
END $$;