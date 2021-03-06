/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */
import type {AlertConfig} from './AlarmAPIType';

import Button from '@fbcnms/ui/components/design-system/Button';
import Grid from '@material-ui/core/Grid';
import HelpIcon from '@material-ui/icons/Help';
import IconButton from '@material-ui/core/IconButton';
import React from 'react';
import Text from '@fbcnms/ui/components/design-system/Text';
import TextField from '@material-ui/core/TextField';
import Tooltip from '@material-ui/core/Tooltip';
import alertsTheme from '@fbcnms/ui/theme/alerts';

import {makeStyles} from '@material-ui/styles';

type Props = {
  alertConfig: AlertConfig,
  setAlertConfig: ((AlertConfig => AlertConfig) | AlertConfig) => void,
  onNext: () => void,
};

const useStyles = makeStyles({
  body: alertsTheme.formBody,
});

export default function AddEditAlertConfigurationStep(props: Props) {
  const classes = useStyles();

  return (
    <>
      <Text variant="h6">CREATE YOUR NEW ALERT</Text>
      <div className={classes.body}>
        <div>
          <Text variant="subtitle1">Configure the alert</Text>
          <Grid container spacing={1} alignItems="flex-end">
            <Grid item>
              <TextField
                required
                placeholder="Ex: up == 0"
                label="Required"
                value={props.alertConfig.expr}
                onChange={event =>
                  props.setAlertConfig({
                    ...props.alertConfig,
                    expr: event.target.value,
                  })
                }
              />
            </Grid>
            <Grid item>
              <Tooltip
                title={
                  'To learn more about how to write alert expressions, click ' +
                  'on the help icon to open the prometheus querying basics guide.'
                }
                placement="right">
                <IconButton
                  href="https://prometheus.io/docs/prometheus/latest/querying/basics/"
                  target="_blank">
                  <HelpIcon />
                </IconButton>
              </Tooltip>
            </Grid>
          </Grid>
        </div>
        <Button onClick={() => props.onNext()}>Next</Button>
      </div>
    </>
  );
}
