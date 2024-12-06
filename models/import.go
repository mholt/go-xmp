// Copyright (c) 2017-2018 Alexander Eichhorn
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package models

// register all metadata models
import (
	_ "github.com/mholt/go-xmp/models/cc"
	_ "github.com/mholt/go-xmp/models/crs"
	_ "github.com/mholt/go-xmp/models/dc"
	_ "github.com/mholt/go-xmp/models/dji"
	_ "github.com/mholt/go-xmp/models/exif"
	_ "github.com/mholt/go-xmp/models/id3"
	_ "github.com/mholt/go-xmp/models/itunes"
	_ "github.com/mholt/go-xmp/models/ixml"
	_ "github.com/mholt/go-xmp/models/mp4"
	_ "github.com/mholt/go-xmp/models/pdf"
	_ "github.com/mholt/go-xmp/models/pm"
	_ "github.com/mholt/go-xmp/models/ps"
	_ "github.com/mholt/go-xmp/models/qt"
	_ "github.com/mholt/go-xmp/models/riff"
	_ "github.com/mholt/go-xmp/models/tiff"
	_ "github.com/mholt/go-xmp/models/xmp_base"
	_ "github.com/mholt/go-xmp/models/xmp_bj"
	_ "github.com/mholt/go-xmp/models/xmp_dm"
	_ "github.com/mholt/go-xmp/models/xmp_mm"
	_ "github.com/mholt/go-xmp/models/xmp_rights"
	_ "github.com/mholt/go-xmp/models/xmp_tpg"
)
