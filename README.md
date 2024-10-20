# GeoGoraphy

- **20-10-2024** :
    1. Made a backend which can recommend user n points around any location based on geospatial data, implemented as a quad tree.
    2. Creates a randomly generated database on it's own and uses that to track the nearest neighbours till either all vertex are considered or all points are achieved.
-  [ **TODO** ] :
    1. Use segment tree to count nodes in a range faster.
    2. Replicate the database as a single database cannot possibly handle all the reads, although the writes can be handled as they usually don't happen that frequently.
