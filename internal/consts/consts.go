// Copyright 2017 Hajime Hoshi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package consts

type Version int

const (
	Version2_5      Version = 0
	VersionReserved Version = 1
	Version2        Version = 2
	Version1        Version = 3
)

type Layer int

const (
	LayerReserved Layer = 0
	Layer3        Layer = 1
	Layer2        Layer = 2
	Layer1        Layer = 3
)

type Mode int

const (
	ModeStereo Mode = iota
	ModeJointStereo
	ModeDualChannel
	ModeSingleChannel
)

const (
	SamplesPerGr  = 576
	BytesPerFrame = SamplesPerGr * 2 * 4
)