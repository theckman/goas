// Tideland Go Application Support - Scroller
//
// Copyright (C) 2014 Frank Mueller / Tideland / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

// Very often there is data continuously written line by line
// which has to be analyzed over the whole time, e.g. the
// monitoring of log files. The Scroller working in the background
// allows to read out of any ReadSeeker (which may be a File) from
// beginning, end or a given number of lines before the end, filter
// the output by a filter function and write it into a Writer. If
// a number of lines and a filter are passed the Scroller tries to
// find that number of lines matching to the filter.
package scroller

// EOF
