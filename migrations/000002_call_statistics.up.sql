CREATE OR REPLACE FUNCTION public.get_call_statistics()
 RETURNS TABLE(chat_id text, call_count integer, latest_call_date timestamptz)
 LANGUAGE plpgsql
 STABLE COST 500
AS $function$
BEGIN
    RETURN QUERY
    SELECT
        c.chat_id,
        COUNT(*)::integer AS call_count,
        MAX(c.created_at) AS latest_call_date
    FROM
        calls c
    GROUP BY
        c.chat_id
    ORDER BY
        call_count DESC
    LIMIT 50;
END;
$function$
;

ALTER TABLE public.excluded ADD CONSTRAINT excluded_item_unique UNIQUE (item);

SELECT setval('excluded_id_seq', COALESCE((SELECT MAX(id) FROM excluded), 1), (SELECT MAX(id) FROM excluded) IS NOT NULL);
