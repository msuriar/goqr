/*
 * Copyright 2015 Murali Suriar
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License in the LICENSE file, or at:
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import "bufio"
import "flag"
import "io/ioutil"
import "os"

import "code.google.com/p/rsc/qr"

func main() {

	out := flag.String("output", "/tmp/qr.png", "Location to output image to.")
	flag.Parse()

	r := bufio.NewReader(os.Stdin)
	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}

	code, err := qr.Encode(string(bytes), qr.H)
	if err != nil {
		panic(err)
	}

	png := code.PNG()

	err = ioutil.WriteFile(*out, png, 0644)
	if err != nil {
		panic(err)
	}
}
