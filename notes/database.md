### accounts
 - `_id` (bson.ObjectId)
 - `handle` (string, e.g. "GeertJohan" in "@GeertJohan", **indexed**)
 - `email` (string, optional)
 - `avatar` (to be decided, link to GridFS file?)

### documents
technical parameters (for system)
 - `_id` (bson.ObjectId)
 - `original` (string, refers to location of the orginal docuement in GridFS)
 - `published` boolean; true: document is visible for users; false: new or not yet processed document
 - `uploaded_date` (time.Time); date of *publication* on Nulpunt. 
content parameters (for people)
 - `uploaderHandle` (string, refers to `accounts.handle`)
 - `title` (string)
 - `summary` (string)
 - `source` (string)
 - `tags` ([]string), as found in the tags collection
 - `originalDate` (time.Time) time of publishing by the gov-ment agency or date of FOI-response

### tags
 - `_id` (bson.ObjectId)
 - `tag` (string)

Note: tags have an ObjectId, these are not for referencing in other collections.
Just insert the tag-string into other collections where needed.

### pages
 - `_id` (bson.ObjectId)
 - `documentId` (bson.ObjectId, refers to `documents._id`)
 - `pageNr` (int, page number)
 - `lines` ([][]char-object)
 - `text` (string); the text in the same order as the lines-attribute, use for search/sharing. Contains ocr-errors

#### char-object
 - `x1` (int, left) in pixels
 - `y1` (int, top) in pixels
 - `x2` (int, bottom) in pixels
 - `y2` (int, right) in pixels
 - `c` (string, character)

### annotations
 - `_id` (bson.ObjectId)
 - `annotatorId` (bson.ObjectId, refers to `accounts._id`)
 - `createDate` (time.Time)
 - `annotation` (string)
 - `location` (object)
  - `start` (object)
   - `page` (int)
   - `x` (int)
   - `y` (int)
  - `end` (object)
   - `page` (int)
   - `x` (int)
   - `y` (int)
 - `comments` ([]comment) structure defined below

#### comment
 - `commenterHandle` (string, refers to `accounts.handle`)
 - `createDate` (time.Time)
 - `comment` (string)
 - `comments` ([]comment) *recursion, disabled for first version??*

### uploads
 - `_id` (bson.ObjectId)
 - `uploaderHandle` (string, refers to `accounts.handle`)
 - `original` (**what's this for??**)
 - `filename` (string)
 - `uploadDate` (time.Time)
 - `language` (string); language of the document to help the OCR (default 'nld')
