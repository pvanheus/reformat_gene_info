# assuming database named gene_info that user can write to

# create db schema
psql gene_info < gene_info.sql

GENE_INFO="<full path to reformatted gene_info file>"
GENE_SYNONYMS="<full path to reformatted gene_info file>"

# use CSV with tab delimiters to load gene_info because it has backslashes in that confuse text format
echo "COPY gene_info FROM '$GENE_INFO' WITH (NULL '-', FORMAT 'csv', DELIMITER E'\t', ESCAPE E'\r', QUOTE E'\r');" \
 | psql gene_info

# assume cleaner format for gene_synonyms - this might be wrong, might need csv trick like above
echo "COPY gene_synonyms FROM '$GENE_SYNONYMS' WITH (FORMAT 'text');" |psql gene_info
