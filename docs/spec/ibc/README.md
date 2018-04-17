# Cosmos Inter-Blockchain Communication (IBC) Protocol

## Abstract

This paper specifies the Cosmos Inter-Blockchain Communication (IBC) protocol. The IBC protocol defines a set of semantics for authenticated, strictly-ordered message passing between two blockchains with independent consensus algorithms.  

IBC requires two blockchains with cheaply verifiable rapid finality. The protocol makes no assumptions of block confirmation times or maximum network latency of packet transmissions, and the two consensus algorithms remain completely independent. Each chain maintains a local partial order and inter-chain message sequencing ensures cross-chain linearity. Once the two chains have registered a trust relationship, cryptographically provable packets can be sent between the chains.

The core IBC protocol is payload-agnostic. On top of IBC, developers can implement the semantics of a particular application, enabling users to transfer valuable assets between different blockchains while preserving, under particular security assumptions of the underlying blockchains, the contractual guarantees of the asset in question - such as scarcity and fungibility for a currency or global uniqueness for a digital kitty-cat. 

IBC was first outlined in the [Cosmos Whitepaper](https://github.com/cosmos/cosmos/blob/master/WHITEPAPER.md#inter-blockchain-communication-ibc), and later described in more detail by the [IBC specification paper](https://github.com/cosmos/ibc/raw/master/CosmosIBCSpecification.pdf). This documentation replaces and supersedes both. It explains the requirements and structure of the protocol and provides sufficient detail for both analysis and implementation, including example pseudocode.

## Contents

1.  **[Overview](overview.md)**
    1.  Definitions
    1.  Threat Models
1.  **[Proofs](proofs.md)**
    1.  Establishing a Root of Trust
    1.  Following Block Headers
1.  **[Messaging Queue](queues.md)**
    1.  Merkle Proofs for Queues
    1.  Naming Queues
    1.  Message Contents
    1.  Sending a Packet
    1.  Receipts
    1.  Relay Process
1.  **[Optimizations](optimizations.md)**
    1.  Cleanup
    1.  Timeout
    1.  Handling Byzantine Failures
1.  **[Conclusion](conclusion.md)**
1.  **[Appendices](appendices.md)**
    1. [Appendix A: Encoding Libraries](appendices.md#appendix-a-encoding-libraries)
    1. [Appendix B: IBC Queue Format](appendices.md#appendix-b-ibc-queue-format)
    1. [Appendix C: Merkle Proof Format](appendices.md#appendix-c-merkle-proof-formats)
    1. [Appendix D: Universal IBC Packets](appendices.md#appendix-d-universal-ibc-packets)
    1. [Appendix E: Tendermint Header Proofs](appendices.md#appendix-e-tendermint-header-proofs)