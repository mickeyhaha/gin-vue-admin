package smt

type ChartData struct {
      Categories  []string `json:"categories"`
      Series  []Series `json:"series"`
}

type Series struct {
      Name string `json:"name"`
      Data []float64 `json:"data"`
      StackGroup string `json:"stackGroup"`
}