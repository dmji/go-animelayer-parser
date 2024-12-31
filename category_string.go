// Code generated by "stringer -type=Category -trimprefix=Category -nametransform=snake_case_lower -fromstringgenfn -outputtransform=snake_case_lower"; DO NOT EDIT.

package animelayer

import "strconv"
import "errors"
import 	"encoding/json"


func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CategoryAnime-0]
	_ = x[CategoryAnimeHentai-1]
	_ = x[CategoryManga-2]
	_ = x[CategoryMangaHentai-3]
	_ = x[CategoryMusic-4]
	_ = x[CategoryDorama-5]
}

const _Category_name = "animeanime_hentaimangamanga_hentaimusicdorama"

var _Category_index = [...]uint8{0, 5, 17, 22, 34, 39, 45}

func (i Category) String() string {
	if i < 0 || i >= Category(len(_Category_index)-1) {
		return "Category(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Category_name[_Category_index[i]:_Category_index[i+1]]
}
func CategoryFromString(s string) (Category, error) {
	for i := 0; i < 6; i++ {
		if e := Category(i + 0); s == e.String() {
			return e, nil
		}
	}
	return Category(0), errors.New("cannot deternime Category from string")
}

func (i Category) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

func (i *Category) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return errors.New("category should be a string")
	}

	var err error
	*i, err = CategoryFromString(s)
	return err
}