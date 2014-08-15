Reformat NCBI gene_info file
============================

The NCBI gene_info.gz file (ftp://ftp.ncbi.nlm.nih.gov/gene/DATA/gene_info.gz) describes all known genes.
Some of its columns, however, are lists, not scalar values. This tool reformats the gene_info file and extracts
the synonyms to their own file, so that the resulting files can be loaded into a PostgreSQL database with COPY
commands. Sample script to load the data is in src/load_gene_info.sh.

The synonyms column is not the only one containing list values: at least chromosome does too, and this script
should be modified to reflect that. Perhaps in the future.

