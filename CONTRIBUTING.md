# Contributing

Contributions a very welcome!

 - Be patient. This is a voluntary effort and things may take time.

## Issues

Issues are welcome however we don't have a lot of resources so:

 - Don't look for fixes on prior versions. Even if upgrading is painful please try the latest version of the package and Go before asking for a change.

## Code

- Please contribute code via a pull request
- Use the latest versions of this package and Go
  - Do not try and backport features or Go compatibility
- Code should be idiomatic to the package
  - This package is careful about naming and package layout to emphasize polymorphism.
- Code should be DRY
  - No syntactic sugar - i.e. no _I wrote language/package X style signatures to existing functions_.
- Balance efficiency with cleanliness
  - Don't write WET code unless it is a significant efficiency/feature gain.
  - Don't write container specific functions where it could be cleanly added to sequence.
