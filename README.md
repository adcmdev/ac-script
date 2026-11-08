# AC GIT COMMIT

## Command Structure

```shell
ac commit [message]
```

## Example

```shell
ac commit "Added new feature"
```

### This command will:

- If the project it's a dart project, it will run `dart format .`
- Add a new line to the commit history
- Stage the changes
- Commit the changes with the provided message
- Push the changes to the remote repository

# AC Checkout

## Command Structure

```shell
ac checkout [branch]
```

## Example

```shell
ac checkout feature/new-login
```

### This command will:

- Switch to the given branch if it already exists
- Create the branch and switch to it if it does not exist

# AC Mobile Deploy Local

## Local Version Change

The `ac mobile deploy local RELEASE_TYPE` command is designed to modify the version in your **pubspec.yaml** and iOS
files. The specific type of version change is determined by the flag you use. Below is a visual guide to the available
release types:

### Command Structure

```shell
ac mobile deploy local [flag]
```

### Release Types

1. **Bug Fix Update (`bugfix`)**
    - Use this for small fixes and patches.
   ```shell
   ac mobile deploy local bugfix
   ```

2. **Minor Update (`minor`)**
    - Ideal for minor feature additions or improvements.
   ```shell
   ac mobile deploy local minor
   ```

3. **Major Update (`major`)**
    - Use this for significant changes, like major feature rollouts.
    ```shell
    ac mobile deploy local major
    ```

### Choosing the Right Flag

- **Bug Fixes**: Small, incremental changes or critical fixes.
- **Minor**: Additions and improvements that don't cause major shifts in functionality.
- **Major**: Comprehensive changes that introduce new features or substantial revisions.

By selecting the appropriate flag, you ensure that your versioning reflects the nature and scope of the changes you're
deploying.
