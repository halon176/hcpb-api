-- public.services definition

-- Drop table

-- DROP TABLE public.services;

CREATE TABLE public.services (
	id int2 NOT NULL,
	"name" text NOT NULL,
	CONSTRAINT services_pk PRIMARY KEY (id)
);


-- public."types" definition

-- Drop table

-- DROP TABLE public."types";

CREATE TABLE public."types" (
	id int2 NOT NULL,
	"name" text NOT NULL,
	CONSTRAINT types_pk PRIMARY KEY (id)
);


-- public.calls definition

-- Drop table

-- DROP TABLE public.calls;

CREATE TABLE public.calls (
	id uuid DEFAULT gen_random_uuid() NOT NULL,
	service_id int2 NOT NULL,
	type_id int2 NOT NULL,
	chat_id text NOT NULL,
	coin text NULL,
	created_at timestamptz DEFAULT now() NOT NULL,
	CONSTRAINT calls_pk PRIMARY KEY (id),
	CONSTRAINT calls_services_fk FOREIGN KEY (service_id) REFERENCES public.services(id),
	CONSTRAINT calls_types_fk FOREIGN KEY (type_id) REFERENCES public."types"(id)
);
CREATE INDEX calls_created_at_idx ON public.calls USING btree (created_at);

-- DROP FUNCTION public.get_statistics();

CREATE OR REPLACE FUNCTION public.get_statistics()
 RETURNS TABLE(distinct_callers integer, daily_calls integer, weekly_calls integer, monthly_calls integer, yearly_calls integer, total_calls integer)
 LANGUAGE plpgsql
 STABLE COST 500
AS $function$
BEGIN
    SELECT 
        COUNT(DISTINCT chat_id),
        COUNT(*) FILTER (WHERE created_at >= date_trunc('day', now()) AND created_at < date_trunc('day', now()) + interval '1 day'),
        COUNT(*) FILTER (WHERE created_at >= date_trunc('week', now()) AND created_at < date_trunc('week', now()) + interval '1 week'),
        COUNT(*) FILTER (WHERE created_at >= date_trunc('month', now()) AND created_at < date_trunc('month', now()) + interval '1 month'),
        COUNT(*) FILTER (WHERE created_at >= date_trunc('year', now()) AND created_at < date_trunc('year', now()) + interval '1 year'),
        COUNT(*)
    INTO 
        distinct_callers, daily_calls, weekly_calls, monthly_calls, yearly_calls, total_calls
    FROM 
        calls;

    RETURN NEXT;

END;
$function$
;


-- public.excluded definition

-- Drop table

-- DROP TABLE public.excluded;

CREATE TABLE public.excluded (
	id serial4 NOT NULL,
	item text NOT NULL,
	CONSTRAINT excluded_pk PRIMARY KEY (id)
);
CREATE INDEX excluded_item_idx ON public.excluded (item);

-- DROP FUNCTION public.get_call_statistics();

CREATE OR REPLACE FUNCTION public.get_call_statistics()
 RETURNS TABLE(chat_id text, call_count bigint, latest_call_date timestamp without time zone)
 LANGUAGE plpgsql
 STABLE COST 500
AS $function$
BEGIN
    RETURN QUERY
    SELECT 
        calls.chat_id, 
        COUNT(*) AS call_count, 
        MAX(calls.created_at)::timestamp AS latest_call_date
    FROM calls
    GROUP BY calls.chat_id
    ORDER BY call_count DESC;
END;
$function$
;
