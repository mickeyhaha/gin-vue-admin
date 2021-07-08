package smt

type ChartData struct {
      Categories  []string `json:"categories"`
      Series  interface{} `json:"series"`
}

type Series struct {
      Name string `json:"name"`
      Data []float64 `json:"data"`
      StackGroup string `json:"stackGroup"`
}

type PieSeries struct {
      Name string `json:"name"`
      Data float64 `json:"data"`
      StackGroup string `json:"stackGroup"`
}


type ColumnLineSeries struct {
      Column []Series `json:"column"`
      Line []Series `json:"line"`
}