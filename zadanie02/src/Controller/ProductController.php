<?php

namespace App\Controller;

use App\Entity\Product;
use App\Repository\ProductRepository;
use Doctrine\ORM\EntityManagerInterface;
use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\JsonResponse;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Attribute\Route;

#[Route('/product')]
final class ProductController extends AbstractController
{
    #[Route(name: 'app_product_index', methods: ['GET'])]
    public function index(ProductRepository $productRepository): JsonResponse
    {
        $products = array_map(
            fn (Product $product): array => $this->productToArray($product),
            $productRepository->findAll()
        );

        return $this->json(['products' => $products]);
    }

    #[Route('', name: 'app_product_create', methods: ['POST'])]
    public function create(Request $request, EntityManagerInterface $entityManager): JsonResponse
    {
        try {
            $data = $this->decodeJson($request);
        } catch (\InvalidArgumentException $exception) {
            return $this->json(['error' => $exception->getMessage()], Response::HTTP_BAD_REQUEST);
        }

        $product = new Product();

        $validationError = $this->applyProductData($product, $data, false);
        if ($validationError !== null) {
            return $this->json(['error' => $validationError], Response::HTTP_UNPROCESSABLE_ENTITY);
        }

        $entityManager->persist($product);
        $entityManager->flush();

        return $this->json($this->productToArray($product), Response::HTTP_CREATED);
    }

    #[Route('/{id}', name: 'app_product_show', methods: ['GET'])]
    public function show(int $id, ProductRepository $productRepository): JsonResponse
    {
        $product = $productRepository->find($id);
        if ($product === null) {
            return $this->json(['error' => 'Product not found.'], Response::HTTP_NOT_FOUND);
        }

        return $this->json($this->productToArray($product));
    }

    #[Route('/{id}', name: 'app_product_update', methods: ['PUT', 'PATCH'])]
    public function update(
        int $id,
        Request $request,
        ProductRepository $productRepository,
        EntityManagerInterface $entityManager
    ): JsonResponse
    {
        $product = $productRepository->find($id);
        if ($product === null) {
            return $this->json(['error' => 'Product not found.'], Response::HTTP_NOT_FOUND);
        }

        try {
            $data = $this->decodeJson($request);
        } catch (\InvalidArgumentException $exception) {
            return $this->json(['error' => $exception->getMessage()], Response::HTTP_BAD_REQUEST);
        }

        $validationError = $this->applyProductData($product, $data, $request->isMethod('PATCH'));
        if ($validationError !== null) {
            return $this->json(['error' => $validationError], Response::HTTP_UNPROCESSABLE_ENTITY);
        }

        $entityManager->flush();

        return $this->json($this->productToArray($product));
    }

    #[Route('/{id}', name: 'app_product_delete', methods: ['DELETE'])]
    public function delete(int $id, ProductRepository $productRepository, EntityManagerInterface $entityManager): JsonResponse
    {
        $product = $productRepository->find($id);
        if ($product === null) {
            return $this->json(['error' => 'Product not found.'], Response::HTTP_NOT_FOUND);
        }

        $entityManager->remove($product);
        $entityManager->flush();

        return $this->json(['message' => 'Product deleted.']);
    }

    private function decodeJson(Request $request): array
    {
        $content = trim($request->getContent());
        if ($content === '') {
            throw new \InvalidArgumentException('Request body cannot be empty.');
        }

        try {
            $data = json_decode($content, true, 512, JSON_THROW_ON_ERROR);
        } catch (\JsonException $exception) {
            throw new \InvalidArgumentException('Invalid JSON payload.');
        }

        if (!is_array($data)) {
            throw new \InvalidArgumentException('JSON payload must be an object.');
        }

        return $data;
    }

    private function applyProductData(Product $product, array $data, bool $partialUpdate): ?string
    {
        $nameProvided = array_key_exists('name', $data);
        $priceProvided = array_key_exists('price', $data);

        if (!$partialUpdate && (!$nameProvided || !$priceProvided)) {
            return 'Fields "name" and "price" are required.';
        }

        if ($nameProvided) {
            if (!is_string($data['name']) || trim($data['name']) === '') {
                return 'Field "name" must be a non-empty string.';
            }

            $product->setName(trim($data['name']));
        }

        if ($priceProvided) {
            if (!is_numeric($data['price'])) {
                return 'Field "price" must be a number.';
            }

            $price = (float) $data['price'];
            if ($price < 0) {
                return 'Field "price" must be greater than or equal to 0.';
            }

            $product->setPrice($price);
        }

        return null;
    }

    private function productToArray(Product $product): array
    {
        return [
            'id' => $product->getId(),
            'name' => $product->getName(),
            'price' => $product->getPrice(),
        ];
    }
}
