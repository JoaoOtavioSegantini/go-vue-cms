import { mount } from "@vue/test-utils";
import ViewPage from "@/components/Posts/View.vue";
import { createStore } from "vuex";

describe("ViewPost.vue", () => {
  afterEach(() => {
    jest.restoreAllMocks();
  });

  const mockRoute = {
    params: {
      id: 1,
    },
  };

  const createVuexStore = () => {
    return createStore({
      state() {
        return {
          Posts: {
            onePost: {
              ID: "1",
              title: "Post 1",
              body: "A simple post text body",
              slug: "simple-slug-text",
            },
          },
        };
      },
      mutations: {
        updatePost(state, res) {
          state.onePost = res;
        },
      },
      actions: {
        getPost(ctx, res) {
          ctx.commit("updatePost", res);
        },
      },
    });
  };

  function factory() {
    const store = createVuexStore();

    return mount(ViewPage, {
      global: {
        plugins: [store],
        mocks: {
          $route: mockRoute,
        },
      },
    });
  }

  it("renders props.msg when passed", async () => {
    const msg =
      "AÇÕES Post 1 Visualização de artigo do blogA simple post text bodyEditarRemover";
    const wrapper = factory();

    console.log(wrapper.emitted());

    expect(mockRoute.params.id).toBe(1);
    expect(wrapper.text()).toMatch(msg);
    expect(wrapper.html()).toMatchSnapshot();
  });
});
