import { chocolatini, moscowMule } from '../../assets/exportCocktailIcons';

import styles from './styles/Cocktails.module.scss';

const cocktailImages: Record<string, string> = {
  'Moscow Mule': moscowMule,
  Chocolatini: chocolatini,
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
  },
];

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
            <h4>{cocktail.name}</h4>
          </div>
        ))}
      </section>
    </section>
  );
};

export default Cocktails;
