// Tideland Go Application Support - Scene
//
// Copyright (C) 2014 Frank Mueller / Tideland / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

// By definition a scene is a sequence of continuous action in a play,
// movie, opera, or book. Applications do know these kind of scenes too.
// Here aspects of the action have to passed between the actors, very
// often they are interwoven and depending.
//
// This also happens in software, especially in concurrent software,
// where different parts have to exchange data in a simple and secure
// way. Here the scene package helps. Beside a simple atomic way to
// store and retrieve information together with optional cleanup
// functions it handles timeouts and signalisation of endings.
package scene

// EOF
