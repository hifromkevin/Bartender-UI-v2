import {
  chocolatini,
  moscowMule,
  ginAndTonic,
} from '../../assets/exportCocktailIcons';

import styles from './styles/Cocktails.module.scss';

const cocktailImages: Record<string, string> = {
  'Moscow Mule': moscowMule,
  Chocolatini: chocolatini,
  'Gin and Tonic': ginAndTonic,
};

const getCocktailImage = (cocktailName: string): string =>
  cocktailImages[cocktailName];

const fakeCocktails = [
  {
    id: 1,
    name: 'Moscow Mule',
    description:
      'A classic cocktail made with whiskey, lemon juice, and sugar.',
    glass_type: 'Cocktail',
    ice_type: 'Cubed',
    requires_shaker: true,
    instructions_id: 1,
    difficulty: 'Medium',
  },
  {
    id: 1,
    name: 'Chocolatini',
    description:
      'A rich, luxurious cocktail made with vodka, chocolate liqueur, and cream.',
    glass_type: 'Martini',
    ice_type: null,
    requires_shaker: false,
    instructions_id: 2,
    difficulty: 'Hard',
  },
  {
    id: 1,
    name: 'Gin and Tonic',
    description: 'A refreshing cocktail made with gin, tonic water, and lime.',
    glass_type: 'Cocktail',
    ice_type: 'Cubed',
    requires_shaker: false,
    instructions_id: 1,
    difficulty: 'Easy',
  },
  {
    id: 1,
    name: 'Chocolatini',
    description:
      'A rich, luxurious cocktail made with vodka, chocolate liqueur, and cream.',
    glass_type: 'Martini',
    ice_type: null,
    requires_shaker: false,
    instructions_id: 2,
    difficulty: 'Hard',
  },
  {
    id: 1,
    name: 'Moscow Mule',
    description:
      'A classic cocktail made with whiskey, lemon juice, and sugar.',
    glass_type: 'Cocktail',
    ice_type: 'Cubed',
    requires_shaker: true,
    instructions_id: 1,
    difficulty: 'Hard',
  },
  {
    id: 1,
    name: 'Chocolatini',
    description:
      'A rich, luxurious cocktail made with vodka, chocolate liqueur, and cream.',
    glass_type: 'Martini',
    ice_type: null,
    requires_shaker: false,
    instructions_id: 2,
    difficulty: 'Hard',
  },
];

const setDifficultyColor = (difficulty: string): string => {
  switch (difficulty) {
    case 'Easy':
      return styles.easy;
    case 'Medium':
      return styles.medium;
    case 'Hard':
      return styles.hard;
    default:
      return '';
  }
};

const Cocktails = () => {
  return (
    <section className={styles.cocktails}>
      <h3>Cocktails</h3>
      <section className={styles.cocktailsList}>
        {fakeCocktails.map((cocktail) => (
          <div key={cocktail.id} className={styles.cocktailItem}>
            <img
              src={getCocktailImage(cocktail.name)}
              alt={cocktail.name}
              className={styles.cocktailImage}
            />
            <h4 className={styles.cocktailName}>
              <span
                className={`${styles.difficulty} ${setDifficultyColor(cocktail.difficulty)}`}
              ></span>
              {cocktail.name}
            </h4>
            {/* <h4
              className={`${styles.cocktailName} ${setDifficultyColor(cocktail.difficulty)}`}
            >
              <span
                className={`${styles.difficulty} ${setDifficultyColor(cocktail.difficulty)}`}
              ></span>
              {cocktail.name}
            </h4> */}
          </div>
        ))}
      </section>
    </section>
  );
};

export default Cocktails;
