# den
Used to quickly store and load data to a .json file as a checkpoint. Useful for testing large operations that you do not wish to rerun.

## Usage

To save the result of a large operation to file for reuse:
```
denKey := "DEN_VALUES"

largeOperationResult := ExecuteLargeOperation()

err := den.Save(denKey, allFirstImages)
if err != nil {
    log.Println("error saving:", err)
}
```

To fetch the result from file:
```
denKey := "DEN_VALUES"

largeOperationResult := make([]string, 0)

err := den.Fetch(denKey, &largeOperationResult)
if err != nil {
    log.Println("error fetching:", err)
}
```

To fetch the result from file without pointer (response will be interface{} and will need to be cast):
```
denKey := "DEN_VALUES"

largeOperationResult, err := den.FetchRaw(denKey)
if err != nil {
    log.Println("error fetching:", err)
}
```

To remove the den file:
```
denKey := "DEN_VALUES"

err := den.Release(denKey)
if err != nil {
    log.Println("error fetching:", err)
}
```