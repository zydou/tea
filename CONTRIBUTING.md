# Contribution Guidelines

## Table of Contents

- [Contribution Guidelines](#contribution-guidelines)
  - [Introduction](#introduction)
  - [Bug reports](#bug-reports)
  - [Discuss your design](#discuss-your-design)
  - [Testing redux](#testing-redux)
  - [Code review](#code-review)
  - [Styleguide](#styleguide)
  - [Sign-off your work](#sign-off-your-work)
  - [Release Cycle](#release-cycle)
  - [Maintainers](#maintainers)
  - [Owners](#owners)
  - [Versions](#versions)
  - [Copyright](#copyright)

## Introduction

This document explains how to contribute changes to TEA.
Sensitive security-related issues should be reported to
[security@gitea.io](mailto:security@gitea.io).

For configuring IDE or code editor to develop Gitea see [IDE and code editor configuration](https://github.com/go-gitea/gitea/tree/master/contrib/ide)

## Bug reports

Please search the issues on the issue tracker with a variety of keywords
to ensure your bug is not already reported.

If unique, [open an issue](https://gitea.com/gitea/tea/issues/new).

Please write clear, concise instructions so we can reproduce the behavior—
even if it seems obvious. The more detailed and specific you are,
the faster we can fix the issue. Check out [How to Report Bugs
Effectively](http://www.chiark.greenend.org.uk/~sgtatham/bugs.html).

Please be kind, remember that TEA comes at no cost to you, and you're
getting free help.

## Discuss your design

The project welcomes submissions. If you want to change or add something,
please let everyone know what you're working on—[file an issue](https://gitea.com/gitea/tea/issues/new)!
Significant changes must go through the change proposal process
before they can be accepted. To create a proposal, file an issue with
your proposed changes documented, and make sure to note in the title
of the issue that it is a proposal.

This process gives everyone a chance to validate the design, helps
prevent duplication of effort, and ensures that the idea fits inside
the goals for the project and tools. It also checks that the design is
sound before code is written; the code review tool is not the place for
high-level discussions.

## Testing redux

Before sending code out for review, run all the test by executing: `make test`
Since TEA is an cli tool it should be obvious to test your feature locally first.

## Code review

Changes to TEA must be reviewed before they are accepted—no matter who
makes the change, even if they are an owner or a maintainer. We use Gitea's
pull request & review workflow to do that. Gitea ensure every PR is reviewed by at least 2 maintainers.

Please try to make your pull request easy to review for us. And, please read
the *[How to get faster PR reviews](https://github.com/kubernetes/community/blob/261cb0fd089b64002c91e8eddceebf032462ccd6/contributors/guide/pull-requests.md#best-practices-for-faster-reviews)* guide;
it has lots of useful tips for any project you may want to contribute.
Some of the key points:

* Make small pull requests. The smaller, the faster to review and the
  more likely it will be merged soon.
* Don't make changes unrelated to your PR. Maybe there are typos on
  some comments, maybe refactoring would be welcome on a function... but
  if that is not related to your PR, please make *another* PR for that.
* Split big pull requests into multiple small ones. An incremental change
  will be faster to review than a huge PR.

## Styleguide

### Commands
- Subcommands should follow the following structure:
    ```
    tea <noun> <verb> [<noun>] [<flags>]
    ```

    for example:

    ```
    tea issues list
    tea pulls create
    tea teams add user --team x --user y
    ```
- Commands should accept nouns as singular & plural by making use of the `Aliases` field.
- The default action without a verb is `list`.
- There is a standard set of verbs: `list/ls`, `create`, `edit`, `delete`
  - `ls` lists objects with filter options, and applies pagination where available.
  - `delete` should show info what is deleted and ask user again, if force flag`-y` is not set
  - Verbs that accept large numbers of flags provide an interactive mode when called without any arguments or flags.
- Try to reuse as many flag definitions as possible, see `cmd/flags/flags.go`.
- Always make sure that the help texts are properly set, and as concise as possible.

### Internal Module Structure
- `cmd`: only contains command/flag options for `urfave/cli`
  - subcommands are in a subpackage named after its parent command
- `modules/task`: contain func for doing something with gitea
  (e.g. create token by user/passwd)
- `modules/print`: contain all functions that print to stdout
- `modules/config`: config tea & login things
- `modules/interact`: contain functions to interact with user by prompts
- `modules/git`: do git related stuff (get info/push/pull/checkout/...)
- `modules/utils`: helper functions used by other functions

### Code Style
Use `make fmt`, check with `make lint`.
For imports you should use the following format (_without_ the comments)
```go
import (
  // stdlib
  "encoding/json"
  "fmt"

  // local packages
  "code.gitea.io/gitea/models"
  "code.gitea.io/sdk/gitea"

  // external packages
  "github.com/foo/bar"
  "gopkg.io/baz.v1"
)
```

## Sign-off your work

The sign-off is a simple line at the end of the explanation for the
patch. Your signature certifies that you wrote the patch or otherwise
have the right to pass it on as an open-source patch. The rules are
pretty simple: If you can certify [DCO](DCO), then you just add a line
to every git commit message:

```
Signed-off-by: Joe Smith <joe.smith@email.com>
```

Please use your real name; we really dislike pseudonyms or anonymous
contributions. We are in the open-source world without secrets. If you
set your `user.name` and `user.email` git configs, you can sign-off your
commit automatically with `git commit -s`.

## Release Cycle

Before we reach v1 there is no fixed release cycle.

## Maintainers

To make sure every PR is checked, we have [team
maintainers](MAINTAINERS). Every PR **MUST** be reviewed by at least
two maintainers (or owners) before it can get merged. A maintainer
should be a contributor of Gitea (or Gogs) and contributed at least
4 accepted PRs. A contributor should apply as a maintainer in the
[Discord](https://discord.gg/Gitea) #develop channel. The owners
or the team maintainers may invite the contributor. A maintainer
should spend some time on code reviews. If a maintainer has no
time to do that, they should apply to leave the maintainers team
and we will give them the honor of being a member of the [advisors
team](https://github.com/orgs/go-gitea/teams/advisors). Of course, if
an advisor has time to code review, we will gladly welcome them back
to the maintainers team. If a maintainer is inactive for more than 3
months and forgets to leave the maintainers team, the owners may move
him or her from the maintainers team to the advisors team.
For security reasons, Maintainers should use 2FA for their accounts and
if possible provide gpg signed commits.
https://help.github.com/articles/securing-your-account-with-two-factor-authentication-2fa/
https://help.github.com/articles/signing-commits-with-gpg/

## Owners

This repo is part of the Gitea project and as such part of that project's
governance.

Since Gitea is a pure community organization without any company support,
to keep the development healthy we will elect three owners every year. All
contributors may vote to elect up to three candidates, one of which will
be the main owner, and the other two the assistant owners. When the new
owners have been elected, the old owners will give up ownership to the
newly elected owners. If an owner is unable to do so, the other owners
will assist in ceding ownership to the newly elected owners.
For security reasons, Owners or any account with write access (like a bot)
must use 2FA.
https://help.github.com/articles/securing-your-account-with-two-factor-authentication-2fa/

After the election, the new owners should proactively agree
with our [CONTRIBUTING](CONTRIBUTING.md) requirements in the
[Discord](https://discord.gg/Gitea) #general channel. Below are the
words to speak:

```
I'm honored to having been elected an owner of Gitea, I agree with
[CONTRIBUTING](CONTRIBUTING.md). I will spend part of my time on Gitea
and lead the development of Gitea.
```

## Versions

tea has the `master` branch as a tip branch and has version branches
such as `release/v0.9`. `release/v0.9` is a release branch and we will
tag `v0.9.0` for binary download. If `v0.9.0` has bugs, we will accept
pull requests on the `release/v0.9` branch and publish a `v0.9.1` tag,
after bringing the bug fix also to the master branch.

Since the `master` branch is a tip version, if you wish to use Gitea
in production, please download the latest release tag version. All the
branches will be protected via GitHub, all the PRs to every branch must
be reviewed by two maintainers and must pass the automatic tests.

## Copyright

Code that you contribute should use the standard copyright header:

```
// Copyright 2020 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

```

Files in the repository contain copyright from the year they are added
to the year they are last changed. If the copyright author is changed,
just paste the header below the old one.
