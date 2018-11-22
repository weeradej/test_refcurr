create or replace procedure SP_DEMOG(PRODUCT_C OUT SYS_REFCURSOR) is
begin

  -- executable part starts here

  OPEN PRODUCT_C FOR
    SELECT * FROM PRODUCT;

  -- EXCEPTION -- exception-handling part starts here
end;
/
