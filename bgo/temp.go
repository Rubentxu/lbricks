(function() {
"use strict";

/
b3.BaseNode = b3.Class(null, {


id: null,
name: null,
category: null,
title: null,
description: null,
parameters: null,
properties: null,

/**
 * Initialization method.
 * @method initialize
 * @constructor
**/
initialize: function(params) {
this.id          = b3.createUUID();
this.title       = this.title || this.name;
this.description = '';
this.parameters  = {};
this.properties  = {};
},


_execute: function(tick) {
// ENTER
this._enter(tick);

// OPEN
if (!tick.blackboard.get('isOpen', tick.tree.id, this.id)) {
this._open(tick);
}

// TICK
var status = this._tick(tick);

// CLOSE
if (status !== b3.RUNNING) {
this._close(tick);
}

// EXIT
this._exit(tick);

return status;
},


_enter: function(tick) {
tick._enterNode(this);
this.enter(tick);
},


_open: function(tick) {
tick._openNode(this);
tick.blackboard.set('isOpen', true, tick.tree.id, this.id);
this.open(tick);
},


_tick: function(tick) {
tick._tickNode(this);
return this.tick(tick);
},


_close: function(tick) {
tick._closeNode(this);
tick.blackboard.set('isOpen', false, tick.tree.id, this.id);
this.close(tick);
},


_exit: function(tick) {
tick._exitNode(this);
this.exit(tick);
},


enter: function(tick) {},


open: function(tick) {},


tick: function(tick) {},


close: function(tick) {},


exit: function(tick) {},
});

})();