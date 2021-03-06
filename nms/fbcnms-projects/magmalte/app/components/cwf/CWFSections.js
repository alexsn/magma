/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */
import type {SectionsConfigs} from '@fbcnms/magmalte/app/components/layout/Section';

import CWFConfigure from './CWFConfigure';
import CWFGateways from './CWFGateways';
import CellWifiIcon from '@material-ui/icons/CellWifi';
import React from 'react';
import SettingsCellIcon from '@material-ui/icons/SettingsCell';

export function getCWFSections(): SectionsConfigs {
  const sections = [
    {
      path: 'gateways',
      label: 'Gateways',
      icon: <CellWifiIcon />,
      component: CWFGateways,
    },
    {
      path: 'configure',
      label: 'Configure',
      icon: <SettingsCellIcon />,
      component: CWFConfigure,
    },
  ];

  return [
    'gateways', // landing path
    sections,
  ];
}
