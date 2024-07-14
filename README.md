
We'll cover three main ideas:

- What is an ORM and why should you be using one?
- How to use GORM
- How does GORM compare to other popular ORMs
  	- Django ORM

---

## What is an ORM?

**Object–relational mapping** is a technique for converting between relational databases and object-oriented programming languages.

| Database | Code |
| --- | --- |
| Table | Class / Model |
| Column | Field / Attribute |
| Relationships | Field / Attribute |

---

## Why use an ORM?

- Typed interaction with your database
- Speed up development time
- Prevent SQL-injection
- Auto migrations
- ➖ Higher level abstraction
- ➖ Can lead to poor performance (N+1)

---

## Relationships

- has one
- has many
- many to many
- polymorphism
- single-table inheritance

---

## Demo — Cars

![](./docs/mermaid-1-simple.svg)

---

## Active record vs. Data mapper

---

## Resources
