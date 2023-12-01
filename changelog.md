# Change Log
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/)
and this project adheres to [Semantic Versioning](http://semver.org/).

## [1.0.1] - 2023-__-__

Description.

### Added
- GET `/v1/folders/:id/images` was created which can be used for pagination. `Limit` and `Page` queries can be added for use in returning specific quantities of data.

### Changed
- Removed media display from the GET `/v1/folders` endpoint and a new GET `/v1/folders/:id/images` was added which returns ONLY a list of images

### Fixed
- Reports generated are now able to display media of different sizes (gb, mb, kb, b)

## [1.0.0] - 2023-11-29

The first production version has been released! Many of the bugs have been tested but its not 100% bug free.
Basic functionality is complete, images and videos are stored, cached, and duplicates are used to save disk storage.
Inputs are validated but has been tested extensively.

### Added

### Changed

### Fixed
