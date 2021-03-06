// @flow
import * as Container from '../../../util/container'
import * as Constants from '../../../constants/wallets'
import * as Types from '../../../constants/types/wallets'
import * as WalletsGen from '../../../actions/wallets-gen'
import * as RouteTreeGen from '../../../actions/route-tree-gen'
import {SendButton as _SendButton, DropdownButton as _DropdownButton} from '.'

type SendButtonOwnProps = {|
  small?: boolean,
|}

const mapStateToPropsSendButton = state => {
  const _account = Constants.getSelectedAccountData(state)
  return {
    _account,
    thisDeviceIsLockedOut: Constants.getThisDeviceIsLockedOut(state, _account.accountID),
  }
}

const mapDispatchToPropsSendButton = dispatch => ({
  _onGoToSendReceive: (from: Types.AccountID, recipientType: Types.CounterpartyType) => {
    dispatch(
      WalletsGen.createOpenSendRequestForm({
        from,
        recipientType,
      })
    )
  },
})

const mergePropsSendButton = (stateProps, dispatchProps, ownProps) => ({
  disabled: !stateProps._account.name || stateProps.thisDeviceIsLockedOut,
  onSendToAnotherAccount: () =>
    dispatchProps._onGoToSendReceive(stateProps._account.accountID, 'otherAccount'),
  onSendToKeybaseUser: () => dispatchProps._onGoToSendReceive(stateProps._account.accountID, 'keybaseUser'),
  onSendToStellarAddress: () =>
    dispatchProps._onGoToSendReceive(stateProps._account.accountID, 'stellarPublicKey'),
  small: ownProps.small,
  thisDeviceIsLockedOut: stateProps.thisDeviceIsLockedOut,
})

export const SendButton = Container.namedConnect<SendButtonOwnProps, _, _, _, _>(
  mapStateToPropsSendButton,
  mapDispatchToPropsSendButton,
  mergePropsSendButton,
  'WalletSendButton'
)(_SendButton)

type DropdownButtonOwnProps = {|
  small?: boolean,
|}

const mapStateToPropsDropdownButton = mapStateToPropsSendButton

const mapDispatchToPropsDropdownButton = dispatch => ({
  _onSettings: (accountID: Types.AccountID) =>
    dispatch(
      RouteTreeGen.createNavigateAppend({
        path: [
          {
            props: {accountID},
            selected: 'settings',
          },
        ],
      })
    ),
  _onShowSecretKey: (accountID: Types.AccountID, walletName: ?string) =>
    dispatch(
      RouteTreeGen.createNavigateAppend({
        path: [
          {
            props: {accountID, walletName},
            selected: 'exportSecretKey',
          },
        ],
      })
    ),
})

const mergePropsDropdownButton = (stateProps, dispatchProps, ownProps) => ({
  disabled: !stateProps._account.name,
  onSettings: () => dispatchProps._onSettings(stateProps._account.accountID),
  onShowSecretKey: stateProps.thisDeviceIsLockedOut
    ? null
    : () => dispatchProps._onShowSecretKey(stateProps._account.accountID, stateProps._account.name),
  small: ownProps.small,
})

export const DropdownButton = Container.namedConnect<DropdownButtonOwnProps, _, _, _, _>(
  mapStateToPropsDropdownButton,
  mapDispatchToPropsDropdownButton,
  mergePropsDropdownButton,
  'WalletDropdownButton'
)(_DropdownButton)
