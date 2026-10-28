# Repository agent guidance

## Working agreements

- Read the relevant source and project documentation before making changes.
- Keep changes focused on the requested task and preserve existing conventions.
- Do not run tests unless the user asks for testing or verification.
- Report files changed and any checks performed in the final response.

## Agent roles

- **Plan agent (orchestrator):** For work that can be split into independent parallel tasks, first map dependencies and define bounded task briefs. Delegate those briefs to generic agents, coordinate integration, resolve conflicts, and check that the combined result meets the request. Keep dependent work sequenced and avoid duplicate edits to shared files.
- **Generic agent:** Handle a clearly scoped implementation or analysis task. Inspect the relevant context, make only the assigned changes, and report the result, files touched, and any blockers to the orchestrator.
- **Research agent:** Investigate questions requiring deep research. Prefer primary and authoritative sources, verify time-sensitive claims, distinguish evidence from inference, and return concise findings with direct citations and implications for the task. Do not modify project files unless assigned.

## Delegation

- Use the plan agent to orchestrate genuinely parallel work; keep straightforward, tightly coupled tasks with one agent.
- Give each delegated agent a distinct scope, expected deliverable, and relevant constraints.
- The orchestrator owns integration and the final response. Review delegated changes before treating the task as complete.
