// Copyright © 2023 Rak Laptudirm <rak@laptudirm.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package search

import "laptudirm.com/x/mess/pkg/search/eval"

// score return the static evaluation of the current context's internal
// board. Any changes to the evaluation function should be done here.
func (search *Context) score() eval.Eval {
	return search.evaluator.Accumulate(search.board.SideToMove)
}

// draw returns a randomized draw score to prevent threefold-repetition
// blindness while searching.
func (search *Context) draw() eval.Eval {
	return eval.RandDraw(search.stats.Nodes)
}
