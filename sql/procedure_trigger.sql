-- Procedure
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

-- Trigger
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