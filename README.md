# Backend Database
The backend database will be MongoDB. Each image will be one object, with the file name, hash, tags, and other metadata.

# Database Populator
This will be a continuously running service that will rescan the import directory on an interval,
looking for new files to add to the database.

In the future, there may also be a slower scheduled integrity checker going around and making
sure that all hashes are up to date and noting untracked or missing files.