export async function fetchCard(cardName: string) {
    const res = await fetch(`https://api.scryfall.com/cards/named?fuzzy=${encodeURIComponent(cardName)}`);

    if (!res.ok) {
        throw new Error(`Failed ot fetch card: ${cardName}`);
    }

    const data = await res.json();

    return {
        name: data.name,
        image: data.image_uris?.normal ?? null,
        typeLine: data.type_line,
        oracleText: data.oracle_text,
        manaCost: data.mana_cost,
        setName: data.set_name,
        prices: data.prices,
        scryfallUrl: data.scryfall_uri,
    };
}