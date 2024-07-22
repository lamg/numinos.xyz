# A platform for value discovery and formalization

Author: [Luis Ángel Méndez Gort](mailto:lamg@protonmail.com)

**_Summary_**
The platform aims to provide the means to experiment with tokenization according to individually defined value criteria, so that anyone, regardless of their financial situation, can contribute to insights into value creation and community dynamics.

## A model use case: SharpieCoin

<p align="center">
    <img
        src="docs/img/the_sharpie_edited.png"
        alt="SharpieCoin project"
        style="border-radius: 50%;width: 100px"
    />
</p>

The F# community, while small, is made up of passionate people who enjoy functional programming, language design, and practical solutions.

The SharpieCoin project has the following goals

- Expose the F# community impact and internal dynamics.

- Formalize social interactions in an independent and neutral way so that they can be effectively and efficiently reproduced outside the platforms on which they originally occurred.

- To incentivize social interactions between token holders.

## Token distribution

The distribution of tokens is guided by a set of criteria that identify value in the community and are determined by collective consensus. The following are proposed for discussion

- Influential papers, books, and blog posts related to F#
- GitHub stars, issues and pull requests
- Nuget downloads
- Social media posts that promote F# and healthy interactions among community members

Those responsible for contributions deemed valuable according to the above points will receive an amount of tokens to be determined later by collective consensus.

There's a possibility users start to exchange their tokens for goods, services or other tokens. The rate in which that happens is called _token velocity_. It's a measure of how successful were the criteria in capturing value and promoting community dynamism. A proposed formula for calculating it in a time period (24 hours, a month, etc.) is:

```fsharp
let tokenVelocity totalVolumeTrades tokenCirculatingSupply =
    totalVolumeTrades/tokenCirculatingSupply
```

Unlike other crypto tokens, **SharpieCoin** is not intended to be deflationary, because it should reflect the value creation by a community, which cannot be predicted at the token's launch. Since in Algorand, the asset creation forces to define a number of tokens, this should be large enough to accomodate the distribution needs of the community.

## Why a blockchain

Blockchains aim to be neutral platforms for storing records transparently, where ownership of data remains in the hands of participants. This is in contrast to popular social media platforms, where you don't own the profile and interactions you create. For example, on X/Twitter you cannot take your handle, likes or record of followers with you to another platform if you choose to do so.

While blockchains are associated with speculation these days, the **SharpieCoin** token is not being distributed with the aim of being listed on cryptocurrency exchanges, or to create a group of privileged privileged insiders who can get rich on empty promises.

## Why Algorand

[Algorand][Algorand] is chosen because it seems to be outside the speculative bubble, is dedicated to technical excellence and is currently focused on improving the developer and user experience. However if the community decides so, the token could be issued on [Cardano][Cardano], [Ergo][Ergo] or any other blockchain.

## The platform

The process of minting and distributing **SharpieCoin** tokens in a blockchain according to the points above suggests that that it would be helpful to have a platform that allows

- the criteria for distributing tokens to be changed and experimented with

- automation of token distribution by consuming data from third-party services, or directly introduced by community managers

- users to redeem their tokens by using third-party authentication services such as GitHub or Mastodon

- other communities to benefit from this type of platform, so we can gain insight into how others value contributions and increase their community dynamism

- publication of a community page displaying their statistics, links, news and achievements

- management of the community public page using a dashboard

- the creation of logos for coins with the help of Large Language Models and templates

[Cardano]: https://cardano.org/
[Algorand]: https://algorand.foundation/
[Ergo]: https://ergoplatform.org/
