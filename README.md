# Simple-Bank

## Database Design

>https://dbdiagram.io/d/63b17f987d39e42284e855e0

```sql
Table accounts as A {
  id bigserial [pk]
  owner varchar [not null]
  balance decimal [not null]
  currency varchar [not null]
  create_at timestamptz [not null, default: `now()`]
  
  indexes {
    owner
  }
}

Table entries {
  id bigserial [pk]
  account_id bigint [not null, ref: > A.id]
  amount decimal [not null, note:'can be negative or positive']
  create_at timestamptz [not null,default: `now()`]
  
  indexes {
    account_id
  }
}

Table transfers {
  id bigserial [pk]
  from_account_id bigint [not null, ref: > A.id]
  to_account_id bigint [not null, ref: > A.id]
  amount decimal [not null,  note:'must be positive']
  create_at timestamptz [not null, default: `now()`]
  
  indexes {
    from_account_id
    to_account_id
    (from_account_id, to_account_id)
  }
}
```
![dbdiagram image](./database/Simple%20Bank.png)
