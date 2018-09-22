# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [v0.7]
### Added
- Function to modify screen mode (fullscreen)
- Variables to modify *Window size and title
- Tests to cover older functions
- More information to README.md to facilitate status and access to some of Macaw data (documentation/lint report)
- Other quit functions in the Macaw.Quit()
- Debug flag so we can silent the console messages from Macaw

### Changed
- Code coverage now uses codecov instead of coverall
- travis.yml now uses test.sh script to execute Macaw packages
- PlaySound() doesn't return chunk anymore
- Changed the way we iterate over the entities
- Render component / system
- Spritesheet and Text files to comply with the new render component / system

### Fixed
- Some minor problems (mispell and lack of comments)

## [v0.6] - 2018-08-10
SHA: 099155a4df8d73b8aab27bbb4d9a01d90cee5a72
### Added
- This changelog :)
- [Wiki page](https://github.com/tubelz/macaw/wiki)
- Issue template on github
- Functions to unload SDL libraries (img, mix, ttf)
- Mock system to help testing
- Function to clear events from a system

### Fixed
- Gameloop test was having problems with the delay time
- Problem with the entity manager deletion algorithm
- Input test failing due to process time

### Changed
- README.md to be more user friendly (badges, lingo)
- Macaw Initialize() now require less arguments

### Removed
- Unused variable from gameloop

## [v0.5] - 2018-05-13
SHA: 30f1d89ac62500f36d2d3ffbf99e2c5fcd5da365
### Added
- Multiple collisions areas per entity
- Grid component
- Types to entities
- Callbacks for Init and Exit scenes
- Test for scenes
- Test for entities
- Test for input package
- Round function for float64

### Fixed
- Input queue

### Removed
- Log from border collision


## [v0.4] - 2018-04-01
SHA: 847761395efe5a93ab62f261e6e08519a553268c
### Added
- Entity manager
- More options for image rendering
- Background color for different scenes
- Camera component
- Internal log to help with tests
- Entity manager
- Gameloop test

### Fixed
- Position according to the camera

### Changed
- Readme.md was updated (dev session)

### Removed
- Setup and Teardown functions in the test for logging purposes
- Skipped frames from gameloop

## [v0.3] - 2018-03-22
SHA: e60e13fcd7ac903350343d12bb2a51ad537c7b1d
### Added
- Mouse event
- Initialization for scenes
- Sound controller with sound options
- Rotation option for images (sprites)
- Camera to render
- Travis.ci 

## [v0.2] - 2018-03-08 
SHA: 3c1bb6ce284c53652644a84cd624d73e6cfb010c 
### Added
- Mouse event
- Scene and SceneManager
- Information on how to install dependencies on README.md
- Test for math package

### Changed
- Initialization
- Some events within physics and collision system

## [v0.1] - 2018-03-04
SHA: a7b2ea6c73522ec5055c9b5a1a90cef534552a27 
### Added
- Initialized project - https://github.com/tubelz/macaw/commit/a7b2ea6c73522ec5055c9b5a1a90cef534552a27

[Unreleased]: https://github.com/tubelz/macaw/compare/v0.6...HEAD
[0.6]: https://github.com/tubelz/macaw/compare/v0.5...v0.6
[0.5]: https://github.com/tubelz/macaw/compare/v0.4...v0.5
[0.4]: https://github.com/tubelz/macaw/compare/v0.3...v0.4
[0.3]: https://github.com/tubelz/macaw/compare/v0.2...v0.3
[0.2]: https://github.com/tubelz/macaw/compare/v0.1...v0.2