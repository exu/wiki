# Quick tips how to do sth [back to index](/)

## `jq`

- show only given fields of json
  ```
  curl https://api.kinguin.net/v1/catalog/regions\?limit\=100 | jq '.[] | "\(.region_id)  \(.name)"'
  ```
  
  
