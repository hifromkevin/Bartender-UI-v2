import React, { useState } from 'react';

import styles from './styles/Tabs.module.scss';

interface ITabData {
  label: string;
  content: string[];
}

const Tabs: React.FC = () => {
  const [activeTab, setActiveTab] = useState(0);

  const tabs: ITabData[] = [
    {
      label: 'Instructions',
      content: [
        'Fill a shaker with ice',
        'Place shaker into bartender',
        'Pour shaker contents into a martini glass',
        'Add a squeeze of lime juice',
        'Add lime wedge as garnish',
      ],
    },
    {
      label: "What You'll Need",
      content: ['Lime Juice', 'Ice', 'Shaker', 'Martini Glass'],
    },
  ];

  return (
    <div className={styles.tabsContainer}>
      <div className={styles.tabsHeader}>
        {tabs.map((tab, index) => (
          <button
            key={index}
            onClick={() => setActiveTab(index)}
            className={`${styles.tabHeading} ${activeTab === index ? styles.active : ''}`}
          >
            {tab.label}
          </button>
        ))}
      </div>
      <div className={styles.tabsContent}>
        <ul>
          {tabs[activeTab]?.content?.map((info: string) => <li>{info}</li>)}
        </ul>
      </div>
    </div>
  );
};

export default Tabs;
