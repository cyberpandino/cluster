/*
 * PandaOS
 * Copyright (C) 2025  Cyberpandino
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU General Public License for more details.
 */

import { proxy } from 'valtio'
import { devtools } from 'valtio/utils'





const initialState: any = {
  session: {
    temparature: {
      min: 0,
      current: 0,
      max: 160,
      ambient: 21,
    },
    rpm: {
      current: 0,
      min: 0,
      max: 8000
    },
    altitude: {
      current: 0
    },
    kilometres: {
      last: 0,
      start: 0,
      current: 138000,
    },
    address: {
      latitude: 0,
      longitude: 0,
    },
    fuel: {
      current: 0,
    },
    speed: {
      current: 0,
      min: 0,
      max: 160
    },
    battery: {
      car: 14.2,
      service: 0.0
    }
  },
  warnings: {
    // Spie standard
    doors: false,
    light: false,
    lowBeam: false,
    highBeam: false,
    fogLight: false,
    engineColant: false,
    warning: false,
    hazard: false,
    turnSignals: false,
    battery: false,
    brakeSystem: false,
    fuel: false,
    injectors: false,
    engineOil: false,
  }

};

const state = proxy(initialState);

export { state }

const unsub = devtools(state, { name: 'state', enabled: import.meta.env.DEV })
