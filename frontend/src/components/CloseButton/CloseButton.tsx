import React from 'react';

import styles from './styles/CloseButton.module.scss';

interface ICloseButton {
  action: Function;
}

const CloseButton: React.FC<ICloseButton> = ({ action }) => {
  return (
    <div className={styles.closeButtonContainer}>
      <button className={styles.closeButton} onClick={() => action(false)}>
        X
      </button>
    </div>
  );
};

export default CloseButton;
