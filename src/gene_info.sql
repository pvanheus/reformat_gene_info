CREATE TABLE gene_synonyms (
    gene_id integer,
    tax_id integer,
    synonym character varying(128)
);

CREATE TABLE gene_info (
    tax_id integer,
    gene_id integer,
    symbol character varying(128),
    locus_tag character varying(32),
    db_xrefs character varying(512),
    chromosome character varying(128),
    map_location character varying(64),
    description character varying(2048),
    type_of_gene character varying(32),
    authority_symbol character varying(32),
    authority_name character varying(256),
    authority_status character varying(4),
    other_designations character varying(4096),
    modification_date date
);