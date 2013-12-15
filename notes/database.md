accounts
 - `_id` (bson.ObjectId)
 - `handle` (string e.g. "GeertJohan" from "@GeertJohan")
 - `email` (string, optional)
 - 

documents
 - `_id` (bson.ObjectId)
 - `accountId` (bson.ObjectId, refers to `accounts._id`)
 - `title` (string)
 - `summary` (string)
 - `source` (string)
 - `categories` ([]string)
 - `publicationDate` (time.Time)
 - `original` (string, refers to location in GridFS)

pages
 - `_id` (bson.ObjectId)
 - `documentId` (bson.ObjectId, refers to `documents._id`)
 - `pageNr` (int, page number)
 - `lines` (two-dimensional array of char-object)

char-object (inside page):
 - `x` (int, left)
 - `y` (int, top)
 - `s` (int, size in pixels)
 - `c` (string, character)
 - 
tags
 - `_id` (bson.ObjectId)
 - `tag` (string)

Note: tags have an ObjectId, these are not for referencing in other tables.
Just insert the tag-string into other tables where needed.

annotations
 - `_id` (bson.ObjectId)
 - `accountId` (bson.ObjectId, refers to `accounts._id`)
 - `createDate` (time.Time)
 - `location` (object)
  - `page` (int)
  - `start` (object)
    - `x` (int)
    - `y` (int)
  - `end` (object)
    - `x` (int)
    - `y` (int)

uploads
 - `_id` (bson.ObjectId)
 - `uploaderId` (bson.ObjectId)
 - `filename` (string)
 - `uploadDate` (time.Time)