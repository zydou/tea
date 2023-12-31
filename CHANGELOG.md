# Changelog

## [v0.9.1](https://gitea.com/gitea/tea/releases/tag/v0.9.1) - 2023-02-15

* BUGFIXES
  * Print pull dont crash if it has TeamReviewRequests (#517)

## [v0.9.0](https://gitea.com/gitea/tea/releases/tag/v0.9.0) - 2022-09-13

* BREAKING
  * Rename master branch to main (#495)
  * Return RFC3339 UTC timestamps for machine-readable output (#470)
* FEATURES
  * Allow editing multiline prompts with external text editor (#429)
  * Add `tea admin user list` (#427)
  * Add `tea whoami` command (#426)
  * Add `tea org create <name>` (#420)
  * Add `tea clone` (#411)
  * Add `tea repo fork` (#410)
  * Add `tea repo create-from-template` (#408)
* BUGFIXES
  * Fetch all items where needed. (#475)
  * Fix running in repos without remote (#472)
  * Add TSV to machine-readable formats (#467)
  * Fix create milestone with deadline bug (#462)
  * Fix resolving of URLs in markdown (#401)
* ENHANCEMENTS
  * Don't emit ANSI sequences when not emitting to TTY for markdown (#491)
  * Show more version info (#486)
  * Add preference `flag_defaults.remote`, refactor (#466)
  * Add `--fields` to notification & milestone listings (#422)
  * PR listing: add --fields & expose additional fields (#415)
  * Add more flags to `tea repo create` (#409)
  * Implement more issue filters (#400)
* MISC
  * Simplify build & update installation instructions (#437)
  * Clarify command descriptions when no arguments are taken (#496)
  * Improve Documentation (#433)
  * Use golang v1.18 and drop vendor folder (#478)
  * Correct spelling of "wether" to "whether" in usage output (#453)

## [v0.8.0](https://gitea.com/gitea/tea/releases/tag/v0.8.0) - 2021-09-22

* BREAKING
  * `tea notifications --all` has moved to `tea notifications --mine` (#389)
  * `tea notifications` now only works with the context of a remote repo. (#389)
    To run this outside of a local git dir, run either tea n `--mine` or `tea n --repo <my/repo>`
* FEATURES
  * Add `tea pr merge` (#348)
* BUGFIXES
  * Don't skip reading the local repo when `--repo` specifies a repo slug (#398)
  * Fix adding login without token on private instances (#392)
  * Correctly match login by ssh host with port (#391)
  * Fix printing issue deadline (#388)
  * Return useful error on wrong sshkey path (#374)
  * Fix parsing of `--description` for issue/pr create (#371)
  * Add missing flags (#369)
  * Check negative limit command parameter (#358) (#359)
  * Add missing flags to org & labels subcommands (#357)
* ENHANCEMENTS
  * Don't require a body for comment PR reviews (#399)
  * Accept more main branch names for login detection (#396)
  * Make local repo optional for `tea pr create`(#393)
  * Notifications Add State Field (#384)
  * Improve error messages (#370)
  * Add tab completion for fish shell (#364)
  * Text editor selection: follow unix defacto standards (#356)
* MISC
  * Update Dependencies (#390)

## [v0.7.1](https://gitea.com/gitea/tea/releases/tag/v0.7.1) - 2021-08-27

* BUILD
  * Enable release builds for darwin/arm64 (#360)

## [v0.7.0](https://gitea.com/gitea/tea/releases/tag/v0.7.0) - 2021-03-12

* BREAKING
  * `tea issue create`: move `-b` flag to `-d` (#331)
  * Drop `tea notif` shorthand in favor of `tea n` (#307)
* FEATURES
  * Add commands for reviews (#315)
  * Add `tea comment` and show comments of issues/pulls (#313)
  * Add interactive mode for `tea milestone create` (#310)
  * Add command to install shell completion (#309)
  * Implement PR closing and reopening (#304)
  * Add interactive mode for `tea issue create` (#302)
* BUGFIXES
  * Introduce workaround for missing pull head sha (#340)
  * Don't exit if we can't find a local repo with a remote matching to a login (#336)
  * Don't push before creating a pull (#334)
  * InitCommand() robustness (#327)
  * `tea comment`: handle piped stdin (#322)
* ENHANCEMENTS
  * Allow checking out PRs with deleted head branch (#341)
  * Markdown renderer: detect terminal width, resolve relative URLs (#332)
  * Add more issue / pr creation parameters (#331)
  * Improve `tea time` (#319)
  * `tea pr checkout`: dont create local branches (#314)
  * Add `tea issues --fields`, allow printing labels (#312)
  * Add more command shorthands (#307)
  * Show PR CI status (#306)
  * Make PR workflow helpers more robust (#300)

## [v0.6.0](https://gitea.com/gitea/tea/releases/tag/v0.6.0) - 2020-12-11

* BREAKING
  * Add `tea repos search`, improve repo listing (#215)
  * Add Detail View for Login (#212)
* FEATURES
  * Add interactive mode for `tea pr create` (#279)
  * Add organization delete command (#270)
  * Add organization list command (#264)
* BUGFIXES
  * Forces needed arguments to `tea ms issues` (#297)
  * Subcommands work outside of git repos (#285)
  * Fix repo flag ignores local repo for login detection (#285)
  * Improve ssh handling (#277)
  * Issue create return web url (#257)
  * Support prerelease gitea instances (#252)
  * Fix `tea pr create` within same repo (#248)
  * Handle login name case-insensitive on all comands (#227)
* ENHANCEMENTS
  * Add `tea login delete` (#296)
  * Release delete: add --delete-tag & --confirm (#286)
  * Sorted milestones list (#281)
  * Pull clean & checkout use token for http(s) auth (#275)
  * Show more infos in pull detail view (#271)
  * Specify fields to print on `tea repos list` (#223)
  * Print times in local timezone (#217)
  * Issue create/edit print details (#214)
  * Improve `tea logout` (#213)
  * Added a shorthand for notifications (#209)
  * Common subcommand naming scheme (#208)
  * `tea pr checkout`: fetch via ssh if available (#192)
  * Major refactor of codebase
* BUILD
  * Use gox to cross-compile (#274)
* DOCS
  * Update Docs to new code structure (#247)

## [v0.5.0](https://gitea.com/gitea/tea/releases/tag/v0.5.0) - 2020-09-27

* BREAKING
  * Add Login Manage Functions (#182)
* FEATURES
  * Add Release Subcomands (#195)
  * Render Markdown and colorize labels table (#181)
  * Add BasicAuth & Interactive for Login (#174)
  * Add milestones subcomands (#149)
* BUGFIXES
  * Fix Pulls Create (#202)
  * Pulls create: detect head branch repo owner (#193)
  * Fix Labels Delete (#180)
* ENHANCEMENTS
  * Add Pagination Options for List Subcomands (#204)
  * Issues/Pulls: Details show State (#196)
  * Make issues & pulls subcommands consistent (#188)
  * Update SDK to v0.13.0 (#179)
  * More Options To Specify Repo (#178)
  * Add Repo Create subcomand & enhancements (#173)
  * Times: format duration as seconds for machine-readable outputs (#168)
  * Add user message to login list view (#166)

## [v0.4.1](https://gitea.com/gitea/tea/releases/tag/v0.4.1) - 2020-09-13

* BUGFIXES
  * Notification don't relay on a repo (#159)

## [v0.4.0](https://gitea.com/gitea/tea/pulls?q=&type=all&state=closed&milestone=1264) - 2020-07-18

* FEATURES
  * Add notifications subcomand (#148)
  * Add subcomand 'pulls create' (#144)
* BUGFIXES
  * Fix Login Detection By Repo Param (#151)
  * Fix Login List Output (#150)
  * Fix --ssh-key Option (#135)
* ENHANCEMENTS
  * Subcomand Login Show List By Default (#152)
* BUILD
  * Migrate src-d/go-git to go-git/go-git (#128)
  * Migrate gitea-sdk to v0.12.0 (#133)
  * Migrate yaml lib (#130)
  * Add gitea-vet (#121)

## [v0.3.1](https://gitea.com/gitea/tea/pulls?q=&type=all&state=closed&milestone=1265) - 2020-06-15

* BUGFIXES
  * --ssh-key should be string not bool (#135) (#137)
  * modules/git: fix dropped error (#127)
  * Issues details: add missing newline (#126)

## [v0.3.0](https://gitea.com/gitea/tea/pulls?q=&type=all&state=closed&milestone=1227) - 2020-04-22

* FEATURES
  * Add `tea pulls [checkout | clean]` commands (#93 #97 #107) (#105)
  * Add `tea open` (#101)
  * Add `tea issues [open|close]` commands (#99)
* ENHANCEMENTS
  * Ignore PRs for `tea issues` (#111)
  * Add --state flag filter to issue & PR lists (#100)

## [v0.2.0](https://gitea.com/gitea/tea/pulls?q=&type=all&state=closed&milestone=538) - 2020-03-06
* FEATURES
  * Add `tea times` command (#54)
* ENHANCEMENTS
  * Upgrade urfave/cli to v2 version (#85)
  * Add --remote flag to add/create subcommands (#77)
* BUILD
  * Upgrade gitea/go-sdk to 2020-01-03 (#81)
  * Update stretchr/testify v1.3.0 -> v1.4.0 (#83)
  * Improve makefile to enable goproxy when go get tools (#98)

## [v0.1.2](https://gitea.com/gitea/tea/pulls?q=&type=all&state=closed&milestone=59) - 2019-11-15
* BUILD
  * Fix typo in drone (#75)

## [v0.1.1](https://gitea.com/gitea/tea/pulls?q=&type=all&state=closed&milestone=59) - 2019-11-15
* FEATURES
  * Add repos subcommand (#65)
* ENHANCEMENTS
  * Minor improvements to command-line language (#66)

## [v0.1.0](https://gitea.com/gitea/tea/pulls?q=&type=all&state=closed&milestone=59) - 2019-10-28
* BREAKING
  * Changed git config determination to go-git (#41) [continue #45] (#62)
* FEATURES
  * Add labels commands (#36)
* BUGFIXES
  * Fix out -o flag (#53)
  * Fix log formatting, refactor flag definition in cmd/labels.go (#52)
* ENHANCEMENTS
  * List label description (#60)
  * Use Different Remote Repos (#58)
  * Unified output (#14) (#40)
  * Added global appendable Flags (#12) (#39)
* BUILD
  * Change .drone.yml to new format (#33)
* DOCS
  * Add install guide from brew on README (#61)
