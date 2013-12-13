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
 - `lines` (two-dimensional array of char-object)

char-object (inside document):
 - `x` (int, left)
 - `y` (int, top)
 - `s` (int, size in pixels)
 - `c` (string, character)

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