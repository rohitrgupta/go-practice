package repeat

func Repeat(s string, count int) string {
  result := ""
  for i := 0; i < count; i ++{
    result = result + s
  }
  return result
}
