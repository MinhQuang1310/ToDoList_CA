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

select * from public.todo_items where title = 'a123456' and status = 'available'
select * from public.todo_items where title = 'a123456'


explain analyze select * from public.todo_items where status = 'available' and title = 'a123456'
explain analyze select * from public.todo_items where title = 'a123456' and status = 'available'

explain analyze select * from public.todo_items where title = 'a123456'

	
drop index idx_todo_items_title_status
drop index idx_todo_items_description
drop index idx_todo_items_title
drop index idx_todo_items_status











