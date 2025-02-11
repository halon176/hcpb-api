CREATE TABLE IF NOT EXISTS public.services (
	id int2 NOT NULL,
	"name" text NOT NULL,
	CONSTRAINT services_pk PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public."types" (
	id int2 NOT NULL,
	"name" text NOT NULL,
	CONSTRAINT types_pk PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.calls (
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

CREATE INDEX IF NOT EXISTS calls_created_at_idx ON public.calls USING btree (created_at);

CREATE TABLE IF NOT EXISTS public.excluded (
	id serial4 NOT NULL,
	item text NOT NULL,
	CONSTRAINT excluded_pk PRIMARY KEY (id)
);
CREATE INDEX excluded_item_idx ON public.excluded (item);

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


INSERT INTO public.types (id, name) VALUES (1, 'Price');
INSERT INTO public.types (id, name) VALUES (2, 'Chart');

INSERT INTO public.services (id, name) VALUES (1, 'CoinGecko');
INSERT INTO public.services (id, name) VALUES (2, 'CoinMarketCap');


