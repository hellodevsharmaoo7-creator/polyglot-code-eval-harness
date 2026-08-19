# Polyglot Code Evaluation Harness 🌐
> **Multi-Language Verification Engine for AI Code Generation & RL Benchmarking**

`polyglot-code-eval-harness` is a unified evaluation framework supporting **Python, TypeScript, Go, Rust, and C++**. Designed for AI training platforms (micro1 RL environment generation), it provides static analysis, AST verification, and automated test execution across diverse programming paradigms.

---

## 💡 Supported Languages & Test Engines

| Language | Test Runner | Benchmarking Tool | Static Analysis / Linter |
| :--- | :--- | :--- | :--- |
| **TypeScript / JS** | Vitest / Jest | Node benchmark | Biome / ESLint |
| **Go** | `go test` | `go test -bench` | `golangci-lint` |
| **Rust** | `cargo test` | `cargo bench` | `cargo clippy` |
| **Python** | `pytest` | `pytest-benchmark` | `ruff` / `mypy` |

---

## 🏗️ Architecture

```
                       +-------------------------------+
                       |  AI Candidate Code Submission |
                       +---------------+---------------+
                                       |
                                       v
                       +---------------+---------------+
                       |   Language Router & Sandbox   |
                       +---------------+---------------+
                                       |
          +--------------------+-------+--------------------+--------------------+
          |                    |                            |                    |
          v                    v                            v                    v
+---------+--------+  +--------+---------+        +---------+--------+  +---------+--------+
| TypeScript Runner|  |    Go Runner     |        |   Rust Runner    |  |  Python Runner   |
| (Vitest / Node)  |  |   (go test)      |        |  (cargo test)    |  |    (pytest)      |
+---------+--------+  +--------+---------+        +---------+--------+  +---------+--------+
          |                    |                            |                    |
          +--------------------+-------+--------------------+--------------------+
                                       |
                                       v
                       +---------------+---------------+
                       | Unified JSON Evaluation Report|
                       +-------------------------------+
```

---

## 📜 License
MIT License. Created for AI Benchmark Training & Polyglot Code Evaluation.
