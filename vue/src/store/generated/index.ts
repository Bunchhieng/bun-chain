// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import BunBun from './bun.bun'
import BunDex from './bun.dex'


export default { 
  BunBun: load(BunBun, 'bun.bun'),
  BunDex: load(BunDex, 'bun.dex'),
  
}


function load(mod, fullns) {
    return function init(store) {        
        if (store.hasModule([fullns])) {
            throw new Error('Duplicate module name detected: '+ fullns)
        }else{
            store.registerModule([fullns], mod)
            store.subscribe((mutation) => {
                if (mutation.type == 'common/env/INITIALIZE_WS_COMPLETE') {
                    store.dispatch(fullns+ '/init', null, {
                        root: true
                    })
                }
            })
        }
    }
}