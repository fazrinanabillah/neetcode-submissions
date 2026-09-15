type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var res strings.Builder
	for _, str := range strs {
		res.WriteString(strconv.Itoa(len(str)))
		res.WriteByte('#')
		res.WriteString(str)
	}
	return res.String()
}

func (s *Solution) Decode(encoded string) []string {
	var res []string
	i := 0
	
	for i < len(encoded) {
		j := i
		
		for encoded[j] != '#' {
			j++
		}

		length, _ := strconv.Atoi(encoded[i:j])
		start := j + 1
		end := start + length
		
		res = append(res, encoded[start:end])
		
		i = end

	}
	return res
}
