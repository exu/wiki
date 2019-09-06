# Quick tips how to do sth

## `jq`

- show only given fields of json
  ```
  curl https://api.kinguin.net/v1/catalog/regions\?limit\=100 | jq '.[] | "\(.region_id)  \(.name)"'
  ```
