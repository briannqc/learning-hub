# GitHub Actions

## Repository - Workflows - Jobs - Steps

1 Repository can have multiple Workflows, listening on different events such as push, fork, issues, issue, pull_request,
release, etc. See
also: [Events that trigger workflows](https://docs.github.com/en/actions/using-workflows/events-that-trigger-workflows).

1 Workflow contains a list of jobs. Jobs run on Runners, e.g., Ubuntu, MacOS, and Windows. They run in parallel by
default and can depend on one another.

1 Job has a list of Steps, e.g., CMD (Typically simple tasks like running command lines), Actions (complex and
repeated tasks), and Docker. Steps in 1 job run in sequence, one after another.

```
Repository 1 ---> N Workflows
Workflow 1 ---> N Jobs
Job 1 ---> N Steps
```

## Workflows - Workflow Runs - Attempts

Workflow is the static definition, executing the Workflow creates a Workflow Run object on GitHub. Workflow Run can
succeed or fail. We can rerun a Workflow Run, which results in another Attempt of such Workflow Run.

```
Workflow (static) 1 ---> N Workflow Runs (dynamic, cancellable, rerunnable)
Workflow Run 1 ---> N Attempts (When rerun)
```

## [Events that trigger workflows](https://docs.github.com/en/actions/using-workflows/events-that-trigger-workflows)

You can configure your workflows to run when specific activity on GitHub happens, at a scheduled time, or when an event
outside of GitHub occurs. Most events also have event activity types to increase the granularity. For
example, Workflows with `on: issues` run for all issues related events, while Workflows with
`on: issues: types: [opened]` only run when new issues are open.

Some common events are:
- push
- issues
  - opened
  - edited
  - deleted
  - transferred
  - pinned
  - unpinned
  - closed
  - reopened
  - assigned
  - unassigned
  - labeled
  - unlabeled
  - locked
  - unlocked
  - milestoned
  - demilestoned
- pull_request:
  - assigned
  - unassigned
  - labeled
  - unlabeled
  - opened
  - edited
  - closed
  - reopened
  - synchronize
  - converted_to_draft
  - locked
  - unlocked
  - enqueued
  - dequeued
  - milestoned
  - demilestoned
  - ready_for_review
  - review_requested
  - review_request_removed
  - auto_merge_enabled
  - auto_merge_disabled
- release
  - published
  - unpublished
  - created
  - edited
  - deleted
  - prereleased
  - released
- workflow_run
  - completed
  - requested
  - in_progress
