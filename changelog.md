# Change Log
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/)
and this project adheres to [Semantic Versioning](http://semver.org/).

## [1.0.1] - 2023-12-03

Changes were made to the endpoints to provided a more user friendly experience. `/folders/:id/images` endpoint can be used to create simple pagination effects when retrieving large amounts of media from folders.
v1.0.0 made use of an extreme amount of concurrency which was unnecessary and hard on the servers processor, which was now removed. The only  concurrency model in place in GIN's default model which allows requests to be handled concurrently at scale.

### Added
- GET `/v1/folders/:id/images` was created which can be used for pagination. `Limit` and `Page` queries can be added for use in returning specific quantities of data.
- Added gzip compression support using the [gin-contrib/gzip](https://github.com/gin-contrib/gzip) module

### Changed
- Removed media display from the GET `/v1/folders` endpoint and a new GET `/v1/folders/:id/images` was added which returns ONLY a list of images

### Fixed
- Reports generated are now able to display media of different sizes (gb, mb, kb, b)
- Removed go routines from the database model. Unnecessary concurrency removed this way, the gin framework handles requests concurrently by default.

## [1.0.0] - 2023-11-29

The first production version has been released! Many of the bugs have been tested but its not 100% bug free.
Basic functionality is complete, images and videos are stored, cached, and duplicates are used to save disk storage.
Inputs are validated but has been tested extensively.

### Added

### Changed

### Fixed
