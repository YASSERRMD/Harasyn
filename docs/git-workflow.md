# Git Workflow

## Harasyn Git Workflow

This project follows a strict phase-based development workflow.

### Git Identity

All commits must be authored as:
```
YASSERRMD <arafath.yasser@gmail.com>
```

Configured via:
```bash
git config user.name "YASSERRMD"
git config user.email "arafath.yasser@gmail.com"
```

### Phase Workflow

For every phase:

```bash
git checkout main
git pull origin main
git checkout -b phase-XX-short-name
```

During the phase:
- Make many atomic commits.
- Validate changes.
- Do not push yet.

Commit message format:
```
phase <number>: <small completed task>
```

At the end of the phase:

```bash
git push -u origin phase-XX-short-name
```

Then:
1. Create PR.
2. Merge PR into `main`.
3. Delete phase branch.
4. Pull latest `main`.
5. Start next phase.

### Author Verification

Before pushing, verify:
```bash
git log --format='%an <%ae>' -n 20
```

Expected result:
```text
YASSERRMD <arafath.yasser@gmail.com>
```

### Branch Naming

- `phase-00-planning`
- `phase-01-foundation`
- `phase-02-data-model`
- etc.

